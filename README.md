# oz_mecab

MeCab 한글 형태소 분석기 Go Wrapper
https://bitbucket.org/eunjeon/mecab-ko-dic

Installation
-----------
* MECAB-KO 형태소 분석기 엔진
```
wget https://bitbucket.org/eunjeon/mecab-ko/downloads/mecab-0.996-ko-0.9.2.tar.gz
tar -xvzf mecab-0.996-ko-0.9.2.tar.gz
cd mecab-0.996-ko-0.9.2
./configure --prefix=/usr/local/mecab
make
make check
make install
```
* MECAB PATH 설정 .bash_profile
```
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin:/usr/local/mecab/bin
```

* MECAB-KO-DIC 사전(DICTIONARY) 파일
```
wget https://bitbucket.org/eunjeon/mecab-ko-dic/downloads/mecab-ko-dic-1.6.1-20140814.tar.gz
tar -xvzf ./mecab-ko-dic-1.6.1-20140814.tar.gz
cd mecab-ko-dic-1.6.1-20140814
./configure
make
make install
```
Test
---------
```
go get github.com/airof98/oz_mecab
```
mecab_test.go
```
package ozmecab

import (
  "github.com/airof98/oz_mecab"
  "fmt"
  "testing"
)   

func TestMecabHelloworld(t *testing.T) {
  e := ozmecab.Init()
  if e != nil {
    t.Error(e)
  }   
  v, e := ozmecab.Parse("정부가 늘어나는 신규 전력수요를 충당하기 위해 2029년까지 원자력발전소")
  if e != nil {
    t.Error(e)
  }   

  defer ozmecab.Fin()
  fmt.Printf("%q\n", v)
}
```
```
go test mecab_test.go -v
```

