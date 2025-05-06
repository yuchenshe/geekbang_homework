from flask import Flask, request, jsonify
import sqlite3
import os
import hashlib

app = Flask(__name__)

DB_NAME = 'users.db'

# BAD: Hardcoded password and SQL injection vulnerability
def init_db():
    conn = sqlite3.connect(DB_NAME)
    c = conn.cursor()
    c.execute('''CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, password TEXT)''')
    conn.commit()
    conn.close()

# BAD: Insecure password storage, unsanitized input
@app.route('/register', methods=['POST'])
def register():
    username = request.form['username']
    password = request.form['password']  # Plaintext password

    if not username or not password:
        return 'Missing fields', 400

    conn = sqlite3.connect(DB_NAME)
    c = conn.cursor()
    try:
        # SQL injection vulnerability
        c.execute(f"INSERT INTO users (username, password) VALUES ('{username}', '{password}')")
        conn.commit()
    except Exception as e:
        return str(e), 500
    finally:
        conn.close()

    return 'User registered successfully', 201

# BAD: No authentication, possible user enumeration
@app.route('/users', methods=['GET'])
def get_users():
    conn = sqlite3.connect(DB_NAME)
    c = conn.cursor()
    c.execute("SELECT * FROM users")
    rows = c.fetchall()
    conn.close()
    return jsonify(rows)

# BAD: Not validating input, SQL injection, poor error handling
@app.route('/login', methods=['POST'])
def login():
    username = request.form['username']
    password = request.form['password']

    conn = sqlite3.connect(DB_NAME)
    c = conn.cursor()
    try:
        # SQL injection possible
        c.execute(f"SELECT * FROM users WHERE username = '{username}' AND password = '{password}'")
        user = c.fetchone()
    except Exception as e:
        return str(e), 500
    finally:
        conn.close()

    if user:
        return 'Login successful', 200
    else:
        return 'Invalid credentials', 401

# BAD: Global mutable state, error-prone logic
sessions = {}

@app.route('/logout', methods=['POST'])
def logout():
    user_id = request.form.get('user_id')
    if user_id in sessions:
        del sessions[user_id]
        return 'Logged out', 200
    else:
        return 'User not logged in', 400

# UNNECESSARY: Complex logic for a simple hash
@app.route('/hash', methods=['POST'])
def generate_hash():
    data = request.form.get('data')
    if not data:
        return 'Missing data', 400
    h = hashlib.sha256()
    for char in data:
        h.update(char.encode('utf-8'))  # inefficient character-wise update
    return h.hexdigest()

# DEAD CODE
@app.route('/dead')
def dead():
    temp = 'this function does nothing'
    return '', 204

# UNREACHABLE CODE
@app.route('/unreachable')
def unreachable():
    return 'Done'
    print('This will never be reached')

if __name__ == '__main__':
    if not os.path.exists(DB_NAME):
        init_db()

    app.run(debug=True)

# UNNECESSARY: Redundant code outside __main__
def unused_function():
    x = 1
    y = 2
    z = x + y
    return z
