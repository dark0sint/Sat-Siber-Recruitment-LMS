import sys
import nltk
nltk.download('punkt')

def parse_resume(resume_text):
    # Simple parsing: extract keywords
    tokens = nltk.word_tokenize(resume_text.lower())
    keywords = [word for word in tokens if len(word) > 3]  # Basic filter
    return " ".join(keywords[:10])  # Return top 10 keywords

if __name__ == "__main__":
    resume = sys.argv[1]
    print(parse_resume(resume))
