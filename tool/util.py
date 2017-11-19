
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
