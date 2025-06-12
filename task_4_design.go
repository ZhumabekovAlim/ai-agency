```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Admin Panel Interface</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        .sidebar {
            width: 250px;
            height: 100vh;
            background-color: #343a40;
            color: #fff;
            position: fixed;
            left: 0;
            top: 0;
            display: flex;
            flex-direction: column;
            padding-top: 10px;
            border-right: 1px solid #ccc;
        }
        .sidebar a {
            color: #fff;
            text-decoration: none;
            padding: 15px;
            border-bottom: 1px solid #444;
        }
        .sidebar a:hover {
            background-color: #495057;
        }
        .content {
            margin-left: 250px;
            padding: 20px;
            height: 100vh;
            overflow-y: auto;
            background-color: #f4f4f9;
        }
        .header {
            padding: 10px;
            background-color: #007bff;
            color: #fff;
            position: sticky;
            top: 0;
        }
        .box {
            background-color: #fff;
            padding: 20px;
            margin-bottom: 20px;
            border-radius: 5px;
            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
        }
        .chat-bot {
            position: fixed;
            bottom: 20px;
            right: 20px;
            width: 300px;
            height: 400px;
            background-color: #fff;
            border: 1px solid #ccc;
            border-radius: 5px;
            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
            display: flex;
            flex-direction: column;
            overflow: hidden;
        }
        .chat-header {
            padding: 10px;
            background-color: #007bff;
            color: #fff;
            text-align: center;
        }
        .chat-messages {
            flex: 1;
            padding: 10px;
            overflow-y: auto;
            background-color: #e9ecef;
        }
        .chat-input {
            padding: 10px;
            display: flex;
            border-top: 1px solid #ccc;
        }
        .chat-input input {
            flex: 1;
            padding: 5px;
            border: 1px solid #ccc;
            border-radius: 3px;
            outline: none;
        }
        .chat-input button {
            padding: 5px 10px;
            background-color: #007bff;
            color: #fff;
            border: none;
            border-radius: 3px;
            cursor: pointer;
            margin-left: 5px;
        }
    </style>
</head>
<body>
    <div class="sidebar">
        <a href="#">Dashboard</a>
        <a href="#">Users</a>
        <a href="#">Settings</a>
        <a href="#">Logs</a>
    </div>
    <div class="content">
        <div class="header">Admin Panel</div>
        <div class="box">Welcome to the Admin Panel. Here you can manage your application.</div>
        <div class="box">This is another interface box for different sections.</div>
    </div>
    <div class="chat-bot">
        <div class="chat-header">Chat with Bot</div>
        <div class="chat-messages"></div>
        <div class="chat-input">
            <input type="text" placeholder="Type a message...">
            <button>Send</button>
        </div>
    </div>
</body>
</html>
```