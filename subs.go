package slug

type subs struct {
	list map[Lang]Sub
}

func newSubs() *subs {
	s := &subs{list: make(map[Lang]Sub)}
	s.Add(TR, TRSub)
	s.Add(EN, ENSub)
	return s
}

func (s *subs) Add(lang Lang, sub Sub) {
	s.list[lang] = sub
}

func (s *subs) Get(lang Lang) Sub {
	return s.list[lang]
}

var langs = newSubs()
