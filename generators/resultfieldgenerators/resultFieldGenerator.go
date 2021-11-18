package resultfieldgenerators

func CreateResultField(
	showRaw bool,
	rawSize int,
	showSnippet bool,
	snipSize int,
	snipFallback bool,
) *SingleResultField {
	var raw *SingleResultFieldRaw
	var snippet *SingleResultFieldSnippet

	if showRaw {
		raw = &SingleResultFieldRaw{}
		if rawSize > 0 {
			raw.Size = rawSize
		}
	}

	if showSnippet {
		snippet = &SingleResultFieldSnippet{}
		if snipSize > 20 {
			snippet.Size = snipSize
		}

		if snipFallback {
			snippet.Fallback = snipFallback
		}
	}

	return &SingleResultField{
		Raw:     raw,
		Snippet: snippet,
	}

}
