package sbd

type Model struct {
	Features map[string]map[string]float64 // ?
}

func NewNBModel(basePath string) {

}

/*
   def classify(self, doc, verbose=False):
       if verbose: sys.stderr.write('NB classifying... ')
       frag = doc.frag
       while frag:
           pred = self.classify_nb_one(frag)
           frag.pred = pred
           frag = frag.next
       if verbose: sys.stderr.write('done!\n')
*/
