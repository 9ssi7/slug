package slug

import "strings"

func new(str string, l ...Lang) string {
	slug := strings.TrimSpace(str)
	sub := langs.Get(getLang(l...))
	for k, v := range sub {
		slug = strings.Replace(slug, string(k), v, -1)
	}
	slug = NonAuthorizedChars.ReplaceAllString(slug, "-")
	slug = MultipleDashes.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-_")
	slug = strings.ToLower(slug)
	return slug
}

func is(slug string) bool {
	if slug == "" ||
		(MaxLength > 0 && len(slug) > MaxLength) ||
		slug[0] == '-' || slug[0] == '_' ||
		slug[len(slug)-1] == '-' || slug[len(slug)-1] == '_' {
		return false
	}
	for _, c := range slug {
		if (c < 'a' || c > 'z') && c != '-' && c != '_' && (c < '0' || c > '9') {
			return false
		}
	}
	return true
}

func addSub(lang Lang, sub Sub) {
	langs.Add(lang, sub)
}

func getLang(l ...Lang) Lang {
	if len(l) > 0 {
		return l[0]
	}
	return EN
}
