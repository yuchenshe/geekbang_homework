from flask import Flask, request, jsonify
import sqlite3
import os
import hashlib

# Initialize Flask application
app = Flask(__name__)

# Database file name
DB_NAME = 'users.db'

def init_db():
    """Initialize the SQLite database and create the users table if it doesn't exist."""
    conn = sqlite3.connect(DB_NAME)
    c = conn.cursor()
    c.execute('''CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, password TEXT)''')
    conn.commit()
    conn.close()

@app.route('/register', methods=['POST'])
def register():
    """Register a new user by inserting username and password into the database."""
    username = request.form['username']
    password = request.form['password']

    if not username or not password:
        return 'Missing fields', 400

    conn = sqlite3.connect(DB_NAME)
    c = conn.cursor()
    try:
        c.execute(f"INSERT INTO users (username, password) VALUES ('{username}', '{password}')")
        conn.commit()
    except Exception as e:
        return str(e), 500
    finally:
        conn.close()

    return 'User registered successfully', 201

@app.route('/users', methods=['GET'])
def get_users():
    """Return the list of all users in the database."""
    conn = sqlite3.connect(DB_NAME)
    c = conn.cursor()
    c.execute("SELECT * FROM users")
    rows = c.fetchall()
    conn.close()
    return jsonify(rows)

@app.route('/login', methods=['POST'])
def login():
    """Authenticate a user based on username and password."""
    username = request.form['username']
    password = request.form['password']

    conn = sqlite3.connect(DB_NAME)
    c = conn.cursor()
    try:
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

# In-memory dictionary for session tracking
sessions = {}

@app.route('/logout', methods=['POST'])
def logout():
    """Log out a user by removing their session entry."""
    user_id = request.form.get('user_id')
    if user_id in sessions:
        del sessions[user_id]
        return 'Logged out', 200
    else:
        return 'User not logged in', 400

@app.route('/hash', methods=['POST'])
def generate_hash():
    """Return SHA-256 hash for a given string input."""
    data = request.form.get('data')
    if not data:
        return 'Missing data', 400
    h = hashlib.sha256()
    for char in data:
        h.update(char.encode('utf-8'))
    return h.hexdigest()

@app.route('/dead')
def dead():
    """Return empty response for a no-op endpoint."""
    temp = 'this function does nothing'
    return '', 204

@app.route('/unreachable')
def unreachable():
    """Endpoint that includes unreachable code after return."""
    return 'Done'
    print('This will never be reached')

if __name__ == '__main__':
    # Create DB file if it doesn't exist
    if not os.path.exists(DB_NAME):
        init_db()

    # Run the Flask development server
    app.run(debug=True)

def unused_function():
    """Simple addition function that is never called."""
    x = 1
    y = 2
    z = x + y
    return z
