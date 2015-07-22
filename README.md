# oz_mecab

MeCab 한글 형태소 분석기 Go Wrapper

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

