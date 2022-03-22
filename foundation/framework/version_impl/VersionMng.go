package version_impl

type VersionMng struct {
	kingdomVersion *Version
	//  kingdomCodeVersion string
	//	kingdomCsvVersion  string
	//	kingdomTomlVersion string

	clientVersion *Version
	//  clientCodeVersion string
	//	clientResVersion  string
}

func NewVersionMng(kversion string) (*VersionMng, error) {
	v, err := NewVersion(kversion)
	if err != nil {
		return nil, err
	}
	return &VersionMng{
		kingdomVersion: v,
	}, nil
}

func (v *VersionMng) SetClientVersion(cversion string) error {
	version, err := NewVersion(cversion)
	if err != nil {
		return err
	}
	v.clientVersion = version
	return nil
}

func (v *VersionMng) CheckKingdomVersion(version string) bool {
	curVersion, err := NewVersion(version)
	if err != nil {
		return false
	}
	b := v.kingdomVersion.GE(*curVersion)
	return b
}

func (v *VersionMng) CheckClientVersion(version string) bool {
	curVersion, err := NewVersion(version)
	if err != nil {
		return false
	}
	b := v.clientVersion.GE(*curVersion)
	return b
}
