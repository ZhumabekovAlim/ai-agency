```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Telegram Bot Interface</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 0;
            padding: 0;
            background-color: #f5f5f5;
        }
        .container {
            max-width: 600px;
            margin: 50px auto;
            padding: 20px;
            background-color: #fff;
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
            border-radius: 8px;
        }
        h1 {
            color: #0088cc;
            text-align: center;
        }
        .input-group {
            margin-bottom: 20px;
        }
        .input-group label {
            display: block;
            margin-bottom: 5px;
            color: #333;
        }
        .input-group input,
        .input-group textarea {
            width: calc(100% - 20px);
            padding: 9px;
            border: 1px solid #ccc;
            border-radius: 4px;
        }
        .input-group textarea {
            resize: vertical;
            min-height: 100px;
        }
        .send-button {
            display: inline-block;
            width: 100%;
            padding: 10px;
            background-color: #0088cc;
            color: #fff;
            border: none;
            border-radius: 4px;
            cursor: pointer;
            text-align: center;
            font-size: 16px;
        }
        .send-button:hover {
            background-color: #007bb5;
        }
    </style>
</head>
<body>

<div class="container">
    <h1>Telegram Bot Interface</h1>
    <form>
        <div class="input-group">
            <label for="message">Message</label>
            <textarea id="message" placeholder="Type your message..."></textarea>
        </div>
        <button type="button" class="send-button">Send</button>
    </form>
</div>

</body>
</html>
```