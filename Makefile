all:
	cgogen -nocgo nl80211.yml

clean:
	rm -f nl80211/const.go

test:
	cd nl80211 && go build
