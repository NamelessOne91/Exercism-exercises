def is_pangram(sentence):
    letters = { letter.lower() for letter in sentence if letter.isalpha()}
    return len(letters) == 26
    
        
