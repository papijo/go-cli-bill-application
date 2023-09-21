# Go CLI Bill Management Application

The Go CLI Bill Management Application is a versatile command-line tool designed to simplify bill creation and management. With this application, you can effortlessly generate bills, add items with their corresponding prices, include tips, and save the finalized bill details to a text file for future reference. Whether you're keeping track of personal expenses, splitting costs with friends, or managing bills for a small business, this application offers an efficient way to handle your financial records.

## Table of Contents

- [Usage](#usage)
- [Files and Structure](#files-and-structure)
- [Features](#features)
- [Contributing](#contributing)
- [License](#license)

## Usage

To utilize the Go CLI Bill Management Application effectively, follow these straightforward steps:

1. **Install Go (if not already installed)**:
   Ensure that you have Go installed on your system. You can download and install it from the official Go website: [https://golang.org/dl/](https://golang.org/dl/).

2. **Download the Application Files**:
   Download the application files (`main.go` and `bill.go`) to a directory of your choice on your local machine.

3. **Navigate to the Application Directory**:
   Open your terminal or command prompt and navigate to the directory where you saved the application files.

4. **Run the Application**:
   Execute the application using the following command:

   ```bash
   go run main.go
   ```

5. **Interact with the Application**:
   Follow the on-screen prompts to interact with the application. You can create a new bill, add items, specify their prices, include tips, and save the bill details.

6. **Save Bill Summary**:
   The application will automatically generate a formatted bill summary. You can choose to save this summary to a text file in the "bills" directory for your records.

## Files and Structure

The Go CLI Bill Management Application is organized into two primary files:

- **`main.go`**: This serves as the main entry point of the application. It manages user input, bill creation, and provides options for adding items, including tips, and saving bills.

- **`bill.go`**: This file defines the `bill` struct and its associated methods. It allows you to create new bills, add items, format bill details, update tips, and save bills to text files.

## Features

### Bill Creation

Creating a new bill is a breeze. Simply provide a name, and the application will automatically generate a bill name with "'s bill" appended to it. This ensures that your bills are organized and labeled consistently.

### Adding Items

Effortlessly add items to your bill by specifying their names and prices. The application supports adding multiple items to a single bill, making it suitable for tracking expenses across various categories.

### Adding Tips

Include tips in your bill with ease. Enter the tip amount in dollars, and the application will intelligently update the total bill, providing you with an accurate summary of your expenses.

### Saving Bills

Preserve your bill details for future reference by saving them to a text file. The application generates a neatly formatted bill summary, and you can choose to save it in the "bills" directory. This feature is particularly useful for maintaining a record of your financial transactions.

## Contributing

We welcome contributions to enhance and extend the functionality of the Go CLI Bill Management Application. If you have ideas for improvements, additional features, or bug fixes, please feel free to contribute by opening issues or submitting pull requests on the project's [GitHub repository](https://github.com/papijo/go-cli-bill-application).

## License

The Go CLI Bill Management Application is licensed under the MIT License.

---

_Note: This application is designed as a versatile tool for managing bills and is provided as-is. It can be customized and expanded to suit your specific needs._
