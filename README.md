# maildigger

little CLI to check for the size of the key of DKIM records.

```


     .-.
    /'v'\           maildigger
   (/   \)           ~djnn.sh
==='="="===<
    |_|              v0.0.1

                                s/o vsim<3
                        hack the planet,
                        travel the world . . .
------------------------------------------------
       DNS scrapping tool to recover DKIM
                   records

    ===> evil.djnn.sh/djnn/maildigger  <===
------------------------------------------------


```

> based on [this research](https://dmarcchecker.app/articles/crack-512-bit-dkim-rsa-key)

## Installing & compiling

Using `go v1.23`.

```bash
# please clone using https. couldnt be arsed to set up a tcp tunnel on cloudflare
git clone https://evil.djnn.sh/djnn/maildigger.git

cd maildigger/

make
```

### Using docker

```bash
cd maildigger/

make docker
```


## Running

```bash
./maildigger -d example.txt --dkim-max-len 128
```
