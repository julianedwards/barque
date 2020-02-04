package rest

import (
	"net/http"

	"github.com/evergreen-ci/barque"
	"github.com/evergreen-ci/barque/model"
	"github.com/evergreen-ci/gimlet"
	"github.com/evergreen-ci/gimlet/ldap"
	"github.com/mongodb/curator/repobuilder"
	"github.com/mongodb/grip"
	"github.com/pkg/errors"
)

type Service struct {
	Environment barque.Environment
	UserManager gimlet.UserManager
	Conf        *model.Configuration
	umconf      gimlet.UserMiddlewareConfiguration
}

func New(env barque.Environment) (*gimlet.APIApp, error) {
	s := &Service{Environment: env}
	if err := s.setup(); err != nil {
		return nil, errors.WithStack(err)
	}
	app := gimlet.NewApp()
	return app, nil
}

func (s *Service) setup() error {
	ctx, cancel := s.Environment.Context()
	defer cancel()
	conf, err := model.FindConfiguration(ctx, s.Environment)
	if err != nil {
		return errors.WithStack(err)
	}
	s.Conf = conf

	if s.Conf.LDAP.URL != "" {
		s.UserManager, err = ldap.NewUserService(ldap.CreationOpts{
			URL:           s.Conf.LDAP.URL,
			Port:          s.Conf.LDAP.Port,
			UserPath:      s.Conf.LDAP.UserPath,
			ServicePath:   s.Conf.LDAP.ServicePath,
			UserGroup:     s.Conf.LDAP.UserGroup,
			ServiceGroup:  s.Conf.LDAP.ServiceGroup,
			PutCache:      model.PutLoginCache,
			GetCache:      model.GetLoginCache,
			ClearCache:    model.ClearLoginCache,
			GetUser:       model.GetUser,
			GetCreateUser: model.GetOrAddUser,
		})
		if err != nil {
			return errors.Wrap(err, "problem setting up ldap user manager")
		}
	} else if s.Conf.NaiveAuth.AppAuth {
		users := []gimlet.User{}
		for _, user := range s.Conf.NaiveAuth.Users {
			users = append(
				users,
				gimlet.NewBasicUser(
					user.ID,
					user.Name,
					user.EmailAddress,
					user.Password,
					user.Key,
					user.AccessRoles,
					user.Invalid,
				),
			)
		}
		s.UserManager, err = gimlet.NewBasicUserManager(users)
		if err != nil {
			return errors.Wrap(err, "problem setting up basic user manager")
		}
	}

	return nil
}

func (s *Service) addRepobuilderJob(rw http.ResponseWriter, r *http.Request) {
	opts := &repobuilder.JobOptions{}
	err := gimlet.GetJSON(r.Body, opts)
	if err != nil {
		panic(err)
	}
	grip.Info(opts)
}
