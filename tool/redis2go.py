import os
import os.path
import glob
import argparse

args = None

def get_file_content(filename, showmsg = True):
    try:
        fp = open(filename, 'rt')
        content = fp.read()
        fp.close()
        return content
    except Exception as e:
        if showmsg:
            print(e, "filename =", filename)
        return ""

def set_file_content(filename, content):
    try:
        fp = open(filename, 'wt')
        fp.write(content)
        fp.close()
        return True
    except Exception as e:
        print(e, "filename =", filename)
        return False

def get_word(content, bs, es, beginpos = 0):
    bi = int(content.find(bs, beginpos)) + len(bs)
    ei = int(content.find(es, bi))
    if ei > 0:
        return content[bi:ei].strip().replace("\n","").replace("\r",""), ei
    else:
        return "", -1

def get_template_file_content(key_type):
    filename = {"s":"template_string", "i":"template_int", "u":"template_uint"}
    content = get_file_content(filename[key_type[0]])
    return content

def dofile(filename, content):
    print("file: ", filename)
    packagename, beginpos = get_word(content, "package", ";")
    classname, beginpos = get_word(content, "message", "\n", beginpos)
    key_type, beginpos = get_word(content, "{", "id", beginpos)
    key_prefix = classname
    print("  ", "packagename:", packagename)
    print("  ", "classname:", classname)
    print("  ", "key_type:", key_type)
    template = get_template_file_content(key_type)
    template = template.replace("{{packagename}}", packagename)
    template = template.replace("{{classname}}", classname)
    template = template.replace("{{key_type}}", key_type)
    template = template.replace("{{key_prefix}}", key_prefix)
    set_file_content("%s/RD_%s.go"%(args.go_out, classname),template)
    print("")

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

