import requests

def main():
    result = requests.get("https://www.googleapis.com/oauth2/v1/certs")
    print(result.text)


if __name__ == "__main__":
    main()
