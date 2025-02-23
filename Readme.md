### **📌 URL Shortening Backend with Go**

## **Introduction**

This is a simple **URL shortening service** built with **Go**. It demonstrates key Go features such as HTTP handling, JSON parsing, concurrency (`sync.Mutex`), and **SHA-256 hashing** for unique URL keys.

🔹 Stores up to **1000 URLs** before automatically resetting.  
🔹 Supports **POST requests** to shorten URLs.  
🔹 Provides redirection via shortened links.

---

## **📖 Usage Guide**

### **🛠 1. Run the Project**

Ensure you have Go installed, then run:

```bash
go run main.go
```

This will start a local server on `http://localhost:8000`.

---

### **📩 2. Shorten a URL**

Send a **POST** request with a JSON body containing the URL you want to shorten.

#### **🔹 Using cURL**

```bash
curl -X POST -H "Content-Type: application/json" -d '{"url":"https://www.example.com"}' http://localhost:8000/shorten
```

#### **🔹 Using Fetch (JavaScript)**

```javascript
const shortenUrl = async (url) => {
  try {
    const response = await fetch("http://localhost:8000/shorten", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ url }),
    });

    if (!response.ok) {
      throw new Error("Failed to fetch the shortened link");
    }

    const data = await response.json();
    console.log("Shortened URL:", data.shortened_url);
    return data.shortened_url;
  } catch (error) {
    console.error(`Error: ${error.message}`);
  }
};

shortenUrl("https://example.com");
```

#### **🔹 Using Axios**

```javascript
const axios = require("axios");

const shortenUrl = async (url) => {
  try {
    const response = await axios.post("http://localhost:8000/shorten", { url });

    console.log("Shortened URL:", response.data.shortened_url);
    return response.data.shortened_url;
  } catch (error) {
    console.error(`Error: ${error.message}`);
  }
};

shortenUrl("https://example.com");
```

---

### **🔄 3. Redirect to Original URL**

Once you have the shortened URL (e.g., `http://localhost:8000/<hashed_key>`), visiting that link will **redirect you** to the original URL.

---

## **📝 Notes**

- The server stores a maximum of **1000 URLs** before resetting the storage.
- URLs are stored **in-memory** and are lost if the server restarts.
- A **hash-based key** is generated using `SHA-256`.
- The `sync.Mutex` ensures thread-safe access to the URL map.

---
