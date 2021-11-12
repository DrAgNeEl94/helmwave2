package plan

import (
	"errors"
	"net/url"
	"os"
	"regexp"

	log "github.com/sirupsen/logrus"
)

var ErrValidateFailed = errors.New("validate failed")

func (p *Plan) ValidateValues() error {
	f := false
	for _, rel := range p.body.Releases {
		for i := range rel.Values {
			_, err := os.Stat(rel.Values[i].Get())
			if os.IsNotExist(err) {
				log.Errorf("❌ %s values (%s): %v", rel.Uniq(), rel.Values[i].Src, err)
				f = true
			} else {
				// FatalError
				return err
			}
		}
	}
	if !f {
		return nil
	}

	return ErrValidateFailed
}

func (p *planBody) Validate() error {
	if len(p.Releases) == 0 && len(p.Repositories) == 0 {
		return errors.New("releases and repositories are empty")
	}

	if err := p.ValidateRepositories(); err != nil {
		return err
	}

	if err := p.ValidateReleases(); err != nil {
		return err
	}

	return nil
}

func (p *planBody) ValidateRepositories() error {
	a := make(map[string]int8)
	for _, r := range p.Repositories {
		if r.Name == "" {
			return errors.New("repository name is empty")
		}

		if r.URL == "" {
			return errors.New("repository url is empty")
		}

		if _, err := url.Parse(r.URL); err != nil {
			return errors.New("cant parse url: " + r.URL)
		}

		a[r.Name]++
		if a[r.Name] > 1 {
			return errors.New("repository name duplicate: " + r.Name)
		}
	}

	return nil
}

func (p *planBody) ValidateReleases() error {
	a := make(map[string]int8)
	for _, r := range p.Releases {
		if r.Name == "" {
			return errors.New("release name is empty")
		}

		if !r.Uniq().Validate() {
			return errors.New("bad uniqname: " + string(r.Uniq()))
		}

		if r.Namespace == "" {
			log.Warnf("namespace for %q is empty. I will use the namespace of your k8s context.", r.Uniq())
		}

		if !validateNS(r.Namespace) {
			return errors.New("bad namespace: " + r.Namespace)
		}

		a[r.Name]++
		if a[r.Name] > 1 {
			return errors.New("release name duplicate: " + r.Name)
		}
	}

	return nil
}

func validateNS(ns string) bool {
	r := regexp.MustCompile("[a-z0-9]([-a-z0-9]*[a-z0-9])?")
	return r.MatchString(ns)
}
