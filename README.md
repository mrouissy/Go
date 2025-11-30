# 📇 Phonebook Application

A feature-rich command-line phonebook application written in Go with JSON persistence.

## ✨ Features

This phonebook application includes **9+ powerful features**:

1. **➕ Add Contact** - Create new contacts with name, phone, email, and address
2. **📋 List All Contacts** - View all your contacts in a formatted list
3. **👁️ View Contact Details** - Display detailed information about a specific contact
4. **✏️ Update Contact** - Modify existing contact information
5. **🗑️ Delete Contact** - Remove contacts with confirmation
6. **🔍 Search Contacts** - Search by name, phone, or email
7. **⭐ Favorite Contacts** - Mark/unmark contacts as favorites
8. **💫 List Favorites** - View only your favorite contacts
9. **📊 Statistics** - See total contacts, favorites, and regular contacts
10. **💾 JSON Persistence** - Automatically save and load data from `phonebook.json`

## 🚀 Getting Started

### Prerequisites

- Go 1.16 or higher installed on your system

### Installation & Running

1. Clone or download this repository
2. Navigate to the project directory:
   ```bash
   cd ~/Desktop/Go
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

### Building an Executable

To compile the application into an executable:

```bash
go build -o phonebook main.go
```

Then run it:
```bash
./phonebook
```

## 📖 Usage

When you run the application, you'll see an interactive menu with numbered options:

```
================================================================================
                          PHONEBOOK APPLICATION
================================================================================

  1.  Add New Contact
  2.  List All Contacts
  3.  View Contact Details
  4.  Update Contact
  5.  Delete Contact
  6.  Search Contacts
  7.  Toggle Favorite
  8.  List Favorite Contacts
  9.  Show Statistics
  0.  Exit
```

### Example Workflow

1. **Add a contact**: Choose option `1` and enter the contact details
2. **Search contacts**: Use option `6` to find contacts by name, phone, or email
3. **Mark as favorite**: Use option `7` and enter the contact ID
4. **View favorites**: Choose option `8` to see all favorite contacts
5. **Exit**: All data is automatically saved to `phonebook.json`

## 💾 Data Storage

All contacts are stored in `phonebook.json` in the same directory as the application. The data includes:

- Contact ID (auto-incremented)
- Name, Phone, Email, Address
- Favorite status
- Creation and update timestamps

The file is automatically created on first run and updated after every modification.

## 🎯 Features in Detail

### CRUD Operations
- **Create**: Add contacts with full details
- **Read**: List all or view individual contacts
- **Update**: Modify any field (leave blank to keep current value)
- **Delete**: Remove contacts with confirmation prompt

### Advanced Features
- **Smart Search**: Case-insensitive search across name, phone, and email fields
- **Favorites System**: Mark important contacts with a ⭐ indicator
- **Statistics**: Track your phonebook usage
- **Auto-save**: Changes are immediately persisted to JSON
- **Timestamps**: Track when contacts were created and last updated

## 🛠️ Technical Details

- **Language**: Go (Golang)
- **Data Format**: JSON
- **Architecture**: Struct-based with methods
- **User Interface**: Interactive CLI with formatted output

## 📝 License

Free to use and modify for personal and educational purposes.

## 👨‍💻 Author

Created as a demonstration of Go programming with file I/O and JSON handling.