// Code generated by ogen, DO NOT EDIT.

package scheme

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

func (s *Server) cutPrefix(path string) (string, bool) {
	prefix := s.cfg.Prefix
	if prefix == "" {
		return path, true
	}
	if !strings.HasPrefix(path, prefix) {
		// Prefix doesn't match.
		return "", false
	}
	// Cut prefix from the path.
	return strings.TrimPrefix(path, prefix), true
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	elemIsEscaped := false
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
			elemIsEscaped = strings.ContainsRune(elem, '%')
		}
	}

	elem, ok := s.cutPrefix(elem)
	if !ok || len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [1]string{}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'd': // Prefix: "dialog/"
				if l := len("dialog/"); len(elem) >= l && elem[0:l] == "dialog/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "user_id"
				// Match until "/"
				idx := strings.IndexByte(elem, '/')
				if idx < 0 {
					idx = len(elem)
				}
				args[0] = elem[:idx]
				elem = elem[idx:]

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'l': // Prefix: "list"
						if l := len("list"); len(elem) >= l && elem[0:l] == "list" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleDialogUserIDListGetRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					case 's': // Prefix: "send"
						if l := len("send"); len(elem) >= l && elem[0:l] == "send" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleDialogUserIDSendPostRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}
					}
				}
			case 'l': // Prefix: "login"
				if l := len("login"); len(elem) >= l && elem[0:l] == "login" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "POST":
						s.handleLoginPostRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "POST")
					}

					return
				}
			case 'u': // Prefix: "user/"
				if l := len("user/"); len(elem) >= l && elem[0:l] == "user/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'g': // Prefix: "get/"
					if l := len("get/"); len(elem) >= l && elem[0:l] == "get/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleUserGetIDGetRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case 'r': // Prefix: "register"
					if l := len("register"); len(elem) >= l && elem[0:l] == "register" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleUserRegisterPostRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}
				case 's': // Prefix: "search"
					if l := len("search"); len(elem) >= l && elem[0:l] == "search" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleUserSearchGetRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				}
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	summary     string
	operationID string
	pathPattern string
	count       int
	args        [1]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// Summary returns OpenAPI summary.
func (r Route) Summary() string {
	return r.summary
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	elem, ok := s.cutPrefix(elem)
	if !ok {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'd': // Prefix: "dialog/"
				if l := len("dialog/"); len(elem) >= l && elem[0:l] == "dialog/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "user_id"
				// Match until "/"
				idx := strings.IndexByte(elem, '/')
				if idx < 0 {
					idx = len(elem)
				}
				args[0] = elem[:idx]
				elem = elem[idx:]

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'l': // Prefix: "list"
						if l := len("list"); len(elem) >= l && elem[0:l] == "list" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: DialogUserIDListGet
								r.name = "DialogUserIDListGet"
								r.summary = ""
								r.operationID = ""
								r.pathPattern = "/dialog/{user_id}/list"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
					case 's': // Prefix: "send"
						if l := len("send"); len(elem) >= l && elem[0:l] == "send" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								// Leaf: DialogUserIDSendPost
								r.name = "DialogUserIDSendPost"
								r.summary = ""
								r.operationID = ""
								r.pathPattern = "/dialog/{user_id}/send"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
					}
				}
			case 'l': // Prefix: "login"
				if l := len("login"); len(elem) >= l && elem[0:l] == "login" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "POST":
						// Leaf: LoginPost
						r.name = "LoginPost"
						r.summary = ""
						r.operationID = ""
						r.pathPattern = "/login"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			case 'u': // Prefix: "user/"
				if l := len("user/"); len(elem) >= l && elem[0:l] == "user/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'g': // Prefix: "get/"
					if l := len("get/"); len(elem) >= l && elem[0:l] == "get/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: UserGetIDGet
							r.name = "UserGetIDGet"
							r.summary = ""
							r.operationID = ""
							r.pathPattern = "/user/get/{id}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}
				case 'r': // Prefix: "register"
					if l := len("register"); len(elem) >= l && elem[0:l] == "register" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							// Leaf: UserRegisterPost
							r.name = "UserRegisterPost"
							r.summary = ""
							r.operationID = ""
							r.pathPattern = "/user/register"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 's': // Prefix: "search"
					if l := len("search"); len(elem) >= l && elem[0:l] == "search" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: UserSearchGet
							r.name = "UserSearchGet"
							r.summary = ""
							r.operationID = ""
							r.pathPattern = "/user/search"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				}
			}
		}
	}
	return r, false
}
