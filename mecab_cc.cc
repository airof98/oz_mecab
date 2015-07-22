
#include <mecab.h>
#include <string.h>
#include <stdlib.h>

extern "C" {
  
  static MeCab::Tagger *tagger = NULL; 
  
  int Init() {
    if (tagger == NULL) {
      tagger = MeCab::createTagger("");

      if (tagger == NULL) {
        fprintf(stderr, "tagger == NULL\n");
        return 0;
      }
    }

    return 1;
  }

  char* Parse(const char* text) {
    if (tagger == NULL) {
      return NULL;
    }

    const char *result = tagger->parse(text);

    if (result == NULL) {
      delete tagger;
      return NULL;
    }

    char* t =  (char*)malloc(strlen(result));
    memcpy(t, result, strlen(result));
    return t;
  }

  void Fin() {
    if (tagger != NULL) {
      delete tagger;
      tagger = NULL;
    }
  }
}