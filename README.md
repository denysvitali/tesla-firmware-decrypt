# tesla-firmware-decrypt

A tool to decrypt the Telsa OTA updates, given a key.

### Requirements

- Go (1.19+)

### Installing

#### From Source

```
git clone https://github.com/denysvitali/tesla-firmware-decrypt
cd tesla-firmware-decrypt
make build install
```

#### From Go

```
go install github.com/denysvitali/tesla-firmware-decrypt/cmd/tesla-firmware-decrypt@latest
```


### Usage

```
tesla-firmware-decrypt \
    -k "uPaonw9KIv1tptQ3hsJsrR2Ibz3YhOcBRBaRitSi4Qg=" \
    your-firmware-file.bin > /tmp/1.squashfs
```