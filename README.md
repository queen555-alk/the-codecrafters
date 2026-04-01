# 🛰️ OPERATION: GOPHER PROTOCOL
## 🚨 MISSION — SENTINEL REBORN

CLASSIFICATION: EYES ONLY
STATUS: THIS IS NOT A DRILL
DURATION: 12 HOURS (09:00 — 21:00)

📡 TRANSMISSION — DIRECTOR X (HQ)

CodeCrafters.

You have come further than anyone expected.

The String Transformer. The Word Surgeon.
The File Pipeline. The Command Console.
One by one — you rebuilt what was broken.

But this mission is different.

Everything you have built so far… were pieces.

Today — you build the machine.

## 🧠 PROJECT OVERVIEW

SENTINEL REBORN is a core text processing engine designed to:

📥 Read raw intelligence (input text/files)  
🧠 Interpret and process structured/unstructured data  
🔄 Apply multiple transformation pipelines  
📤 Output refined, actionable intelligence  

This system simulates a real-world data processing pipeline, combining:  

- Text parsing
- Command-based transformations
- File I/O handling
- Modular function execution
- ⚙️ CORE OBJECTIVE

### Build a modular, extensible Go-based engine that:

- Processes input text dynamically
- Applies transformation commands
- Handles edge cases gracefully
- Writes results to output streams/files

## 🏗️ PROJECT STRUCTURE
sentinel-reborn/  
│  
├── main.go                # Entry point of the application  
│  
└── functions/            # Core transformation modules  
    ├── bin.go            # Binary conversion logic  
    ├── cap.go            # Capitalization logic  
    ├── commandN.go       # Command parsing & execution  
    ├── hex.go            # Hexadecimal conversion logic  
    ├── lower.go          # Lowercase transformation  
    ├── puctuation.go     # Punctuation handling  
    ├── singlequotes.go   # Single quote formatting  
    ├── up.go             # Uppercase transformation  
    └── vowels.go         # Vowel processing rules  

## 🧩 SYSTEM DESIGN

The system follows a pipeline architecture:

- INPUT → PARSE → TRANSFORM → OUTPUT
- 🔄 Transformation Flow
- Input ingestion
- Read from file or stdin
- Command parsing
- Detect transformation instructions
- Execution pipeline
- Apply transformations in sequence:
- Case conversions (upper, lower, capitalize)
- Encoding (hex, bin)
- Formatting (quotes, punctuation)
- Linguistic rules (vowels, grammar)
- Output generation
- Write processed text to stdout or file
- 🧠 KEY FEATURES

## ✅ Modular Function Design

Each transformation is isolated in its own file for:

- Reusability
- Maintainability
- Testability
- ✅ Command-Based Processing

Dynamic commands control behavior:

Examples:

(up) → Convert to uppercase
(low) → Convert to lowercase
(cap) → Capitalize text
(hex) → Convert hex to decimal
(bin) → Convert binary to decimal
✅ Intelligent Text Handling
Handles punctuation spacing
Fixes quote formatting
Applies grammar-aware vowel rules (a → an)
✅ Extensible Architecture

New transformations can be added easily by:

Creating a new function file
Registering it in the command handler
🚀 HOW TO RUN
1. Clone the repository
git clone <your-repo-url>
cd sentinel-reborn
2. Initialize Go module (if needed)
go mod tidy
3. Run the program
go run . input.txt output.txt  


🧪 SAMPLE USAGE
Input:
this is a test (up)
Output:
THIS IS A TEST
Input:
1E (hex)
Output:
30


## 👥 TEAM — DEBUGGERS
Name |	Role  | function implemented  
-----|--------|------------
Victor Akkor |	Code Legend |  
Queen |	Team Lead  |  
Timothy Ododo |	Developer | Single Quote function  
Adah Jerry | Developer | command marker  
Adanu Gabriel Oche | Developer | main function  


## 🧙 THE CODE LEGEND SYSTEM

Each squad is assigned a Code Legend whose role is to:

🧭 Guide architectural decisions
❓ Ask critical thinking questions
🚧 Unblock stuck team members
⏱️ Run checkpoints (1PM & 5PM)
🧩 Assign micro-tasks to maintain momentum

They do not give answers.
They unlock thinking.

## ⏱️ MISSION TIMELINE
Time	Activity
09:00 AM	Mission Briefing
09:30 AM	Planning Complete
01:00 PM	Checkpoint 1
05:00 PM	Checkpoint 2
09:00 PM	Final Assembly
## 🧠 DEVELOPMENT PRINCIPLES
Plan before code  
Write modular functions  
Handle edge cases early  
Test continuously  
Collaborate, then assemble  
## ⚠️ RULES OF ENGAGEMENT
❌ No shortcuts  
❌ No copy-paste solutions  
✅ Think before coding  
✅ Build independently  
✅ Integrate as a team  
🔥 FINAL DIRECTIVE  

SENTINEL will either come back online tonight —
or it will not.

That decision belongs entirely to you.

## 🏁 CONCLUSION

This project is more than just code.

It is:

A test of logic
A test of endurance
A test of teamwork

And most importantly —
a test of what you are capable of building under pressure.

## 📜 LICENSE

This project is developed as part of the Learn2Earn CodeCrafters Program.