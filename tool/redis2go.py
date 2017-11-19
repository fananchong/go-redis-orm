import os
import os.path
import glob
import argparse

args = None


def dofile(filename, content):
    print("file: ", filename)


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='redis2go',formatter_class=argparse.RawDescriptionHelpFormatter)
    parser.add_argument("--proto_path", default="../example/proto", help="proto path", type=str)
    parser.add_argument("--go_out", default="../example", help="out path", type=str)

    args = parser.parse_args()

    for dir, _, _ in os.walk(args.proto_path):
        filenames = glob.glob( dir + '/*.proto')
        for filename in filenames:
            f = open(filename, 'rt')
            content = f.read()
            f.close()
            dofile(filename, content)
    print("done.")

