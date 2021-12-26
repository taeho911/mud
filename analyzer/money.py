import sys

def foo(bar: str) -> str:
    return f'foo :: {bar}'

if __name__ == '__main__':
    print(foo(sys.argv[1]))
    print(foo([1, 2]))
