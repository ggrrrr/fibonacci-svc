package fi

type (
	Number uint64
	Fi     struct {
		Previous Number `json:"previous"`
		Current  Number `json:"current"`
	}
)

func Next(fi Fi) Fi {
	if fi.Current == 0 {
		return Fi{
			Previous: 0,
			Current:  1,
		}
	}
	return Fi{
		Previous: fi.Current,
		Current:  fi.Previous + fi.Current,
	}
}

func Previous(fi Fi) Fi {
	if fi.Current == 0 {
		return Fi{
			Previous: 0,
			Current:  0,
		}
	}
	if fi.Previous == 0 {
		return Fi{
			Previous: 0,
			Current:  0,
		}
	}
	return Fi{
		Current:  fi.Previous,
		Previous: fi.Current - fi.Previous,
	}
}
