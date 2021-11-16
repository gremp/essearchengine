package helpers

func CreateResultSetting(
	showRaw bool,
	rawSize int,
	showSnippet bool,
	snipSize int,
	snipFallback bool,
) *SingleResultSettings {
	var raw *SingleResultSettingsRaw
	var snippet *SingleResultSettingsSnippet

	if showRaw {
		raw = &SingleResultSettingsRaw{}
		if rawSize > 0 {
			raw.Size = rawSize
		}
	}

	if showSnippet {
		snippet = &SingleResultSettingsSnippet{}
		if snipSize > 20 {
			snippet.Size = snipSize
		}

		if snipFallback {
			snippet.Fallback = snipFallback
		}
	}

	return &SingleResultSettings{
		Raw:     raw,
		Snippet: snippet,
	}
}

func CreateSearchSetting(weight int) *SingleFieldSettings {
	return &SingleFieldSettings{
		Weight: weight,
	}
}
