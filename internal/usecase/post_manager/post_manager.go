package post_manager

import (
	"context"

	"github.com/go-faster/errors"

	feedCache "monolith/internal/cache"
	"monolith/internal/models"
)

type PostManager struct {
	storage      storage
	cache        cache
	friendGetter friendGetter
}

func New(storage storage, cache cache, friendGetter friendGetter) *PostManager {
	return &PostManager{
		storage:      storage,
		cache:        cache,
		friendGetter: friendGetter,
	}
}

func (p *PostManager) Create(ctx context.Context, post models.Post) (string, error) {
	postID, err := p.storage.InsertPost(ctx, post)
	if err != nil {
		return "", errors.Wrap(err, "insert post error")
	}

	post.ID = postID
	friends, err := p.friendGetter.GetFriends(ctx, post.UserID)
	if err != nil {
		return "", errors.Wrap(err, "get friend by user id error")
	}
	for _, friend := range friends {
		if err := p.cache.Append(ctx, friend.UserID, post); err != nil {
			return "", errors.Wrap(err, "append post in friend cache error")
		}
	}

	return postID, nil
}

func (p *PostManager) Update(ctx context.Context, post models.Post) error {
	savedPost, err := p.storage.FindByPostID(ctx, post.ID)
	if err != nil {
		return errors.Wrap(err, "find by post id error")
	}

	if savedPost == nil {
		return nil
	}

	if savedPost.UserID != post.UserID {
		return errors.New("user is not post author")
	}

	if err := p.storage.UpdatePost(ctx, post); err != nil {
		return errors.Wrap(err, "update post error")
	}

	friendships, err := p.friendGetter.GetFriends(ctx, post.UserID)
	if err != nil {
		return errors.Wrap(err, "get friend by user id error")
	}

	for _, friendship := range friendships {
		if err := p.cache.Update(ctx, friendship.UserID, post); err != nil {
			return errors.Wrap(err, "update friend cache error")
		}
	}

	return nil
}

func (p *PostManager) GetByID(ctx context.Context, postID string) (*models.Post, error) {
	post, err := p.storage.FindByPostID(ctx, postID)
	if err != nil {
		return nil, errors.Wrap(err, "find by post id error")
	}

	return post, nil
}

func (p *PostManager) Delete(ctx context.Context, post models.Post) error {
	savedPost, err := p.storage.FindByPostID(ctx, post.ID)
	if err != nil {
		return errors.Wrap(err, "find by post id error")
	}

	if savedPost == nil {
		return nil
	}

	if savedPost.UserID != post.UserID {
		return errors.New("user is not post author")
	}

	if err := p.storage.DeletePost(ctx, post.ID); err != nil {
		return errors.Wrap(err, "delete post error")
	}

	friendships, err := p.friendGetter.GetFriends(ctx, post.UserID)
	if err != nil {
		return errors.Wrap(err, "get friend by user id error")
	}

	for _, friendship := range friendships {
		if err := p.cache.Delete(ctx, friendship.UserID, *savedPost); err != nil {
			return errors.Wrap(err, "delete post from friend cache error")
		}
	}

	return nil
}

func (p *PostManager) InitFeedCache(ctx context.Context) error {
	friendships, err := p.friendGetter.GetAllFriends(ctx)
	if err != nil {
		return errors.Wrap(err, "get all friendship error")
	}

	for _, friendship := range friendships {
		if err := p.cache.Clear(ctx, friendship.UserID); err != nil {
			return errors.Wrap(err, "clear cache by user id error")
		}
	}

	for _, friendship := range friendships {
		posts, err := p.storage.FindPostsByUserID(ctx, friendship.FriendID, feedCache.CacheLen)
		if err != nil {
			return errors.Wrap(err, "get posts by friendship id error")
		}

		for _, post := range posts {
			if err := p.cache.Append(ctx, friendship.UserID, post); err != nil {
				return errors.Wrap(err, "append post in friendship cache error")
			}
		}
	}

	return nil
}
