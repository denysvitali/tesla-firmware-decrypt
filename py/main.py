import base64
import tempfile

import click
import salsa20


@click.command()
@click.option("-i", "--input", required=True)
@click.option("-k", "--key", required=True)
def cmd(input: str, key: str):
    tmp = tempfile.mktemp()
    print(tmp)

    with open(tmp, mode="wb") as tmp_out:
        with open(input, mode="rb") as f:
            counter = 0
            key_bytes = base64.b64decode(key)
            while True:
                in_bytes = f.read(256)
                if len(in_bytes) == 0:
                    return
                counter_bytes = counter.to_bytes(24, byteorder='little')
                res = salsa20.XSalsa20_xor(in_bytes, counter_bytes, key_bytes)
                tmp_out.write(res)
                counter += 1

if __name__ == "__main__":
    cmd()
