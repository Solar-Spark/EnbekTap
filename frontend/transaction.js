document
  .getElementById("transactionForm")
  .addEventListener("submit", async function (e) {
    e.preventDefault();

    const card = document.getElementById("cardNumber").value;
    const cvv = document.getElementById("cvv").value;
    const name = document.getElementById("cardName").value;
    const method = document.getElementById("paymentMethod").value;
    const amount = document.getElementById("amount").value;
    const token = localStorage.getItem("access_token");

    // Basic validation
    if (!/^\d{16}$/.test(card)) {
      alert("Please enter a valid 16-digit card number");
      return;
    }

    if (!/^\d{3}$/.test(cvv)) {
      alert("Please enter a valid 3-digit CVV");
      return;
    }

    if (amount <= 0) {
      alert("Please enter a valid amount");
      return;
    }

    // Create transaction object
    const transaction = {
      card,
      cvv,
      name,
      method,
      amount: Number(document.getElementById("amount").value),
    };

    console.log("Sending Data:", JSON.stringify(transaction));

    try {
      const response = await fetch(
        "http://localhost:8080/auth/createtransaction",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
          },
          body: JSON.stringify(transaction),
        }
      );

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data = await response.json();
      console.log("Success:", data);
      alert("Payment submitted successfully!");
      e.target.reset();
      await fetchTransactionHistory();
    } catch (error) {
      console.error("Error:", error);
      alert("Failed to process payment. Please try again.");
    }
  });

// Add input validation for card number (numbers only)
document.getElementById("cardNumber").addEventListener("input", function (e) {
  this.value = this.value.replace(/[^\d]/g, "");
});

// Add input validation for CVV (numbers only)
document.getElementById("cvv").addEventListener("input", function (e) {
  this.value = this.value.replace(/[^\d]/g, "");
});

// Add this function after your existing code

async function fetchTransactionHistory() {
  try {
    const token = localStorage.getItem("access_token");
    const response = await fetch(`http://localhost:8080/auth/transactions`, {
      method: "GET",
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const transactions = await response.json();
    displayTransactions(transactions);
  } catch (error) {
    console.error("Error fetching transactions:", error);
  }
}

function displayTransactions(transactions) {
  const tableBody = document.getElementById("transactionTableBody");
  tableBody.innerHTML = ""; // Clear existing rows

  transactions.forEach((transaction) => {
    const row = document.createElement("tr");
    row.innerHTML = `
            <td>${transaction.transaction_id}</td>
            <td class="status-${transaction.status.toLowerCase()}">${
      transaction.status
    }</td>
        `;
    tableBody.appendChild(row);
  });
}

// Call this function when the page loads
document.addEventListener("DOMContentLoaded", fetchTransactionHistory);
