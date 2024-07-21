def convert_to_lowercase(input_file, output_file):
    try:
            # Open the input file in read mode
        with open(input_file, 'r') as infile:
                # Read all lines from the input file
            lines = infile.readlines()
            
            # Convert each line to lowercase
            lowercase_lines = [line.lower() for line in lines]
        
        # Open the output file in write mode
        with open(output_file, 'w') as outfile:
                # Write the lowercase lines to the output file
            outfile.writelines(lowercase_lines)
            
        print(f"Conversion completed. Lowercase words saved to {output_file}.")
    except FileNotFoundError:
        print(f"The file {input_file} does not exist.")
    except Exception as e:
        print(f"An error occurred: {e}")

# Example usage
input_file = 'wordleWords.txt'
output_file = 'lowercase_words.txt'
convert_to_lowercase(input_file, output_file)

