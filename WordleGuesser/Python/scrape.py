import requests
from lxml import html

def scrape_xpath(url, xpath):
    response = requests.get(url)
    response.raise_for_status()  # Check that the request was successful
    
    # Parse the HTML content
    tree = html.fromstring(response.content)
        
    # Extract values using XPath
    values = tree.xpath(xpath)
    
    # Extract the text content from the elements
    extracted_values = [value.text_content() for value in values]
        
    return extracted_values

def save_to_file(values, filename):
    with open(filename, 'w') as file:
        for value in values:
            file.write(f"{value}\n")

# Example usage
alphabet = ['b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z']
letter = 'z'
pages = 1
n = 1

while n <= pages:
   url = f'https://www.merriam-webster.com/wordfinder/classic/begins/all/5/{letter}/{n}'  # Replace with the actual URL
   filename = f'dictionary/{letter}{n}.txt'
   xpath = '/html/body/div/div/div[2]/div/div[1]/div/div/div[2]/div/div[1]/div/div[2]/div/div[2]/ul'
   values = scrape_xpath(url, xpath)
   save_to_file(values, filename)
   n+=1
   print(f"Values saved to {filename}")
 



