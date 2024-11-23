import React, { useEffect, useState } from 'react';
import './App.css';
import Todo, { TodoType } from './Todo';

function App() {
  const [todos, setTodos] = useState<TodoType[]>([]);
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');

  // Initially fetch todo
  useEffect(() => {
    const fetchTodos = async () => {
      try {
        const todos = await fetch('http://localhost:8080/');
        if (todos.status !== 200) {
          console.log('Error fetching data');
          return;
        }

        setTodos(await todos.json());
      } catch (e) {
        console.log('Could not connect to server. Ensure it is running. ' + e);
      }
    }

    fetchTodos()
  }, []);

  // Handles submitted ToDos
  const submitToDo = async (event: React.FormEvent) => {
    event.preventDefault();
    try {
      const response = await fetch('http://localhost:8080/', {
        method: 'POST',
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          title: title,
          description: description
        }),
      });

      if (!response.ok) {
        throw new Error(`HTTP error, Status ${response.status}`);
      }

      console.log('ToDo Submitted');
    } catch (e) {
      console.log('Could not submit To-Do ' + e);
    }

    // Updates ToDo List
    const todos = await fetch('http://localhost:8080/');
        if (todos.status !== 200) {
          console.log('Error fetching data');
          return;
        }

    setTodos(await todos.json());
  }

  return (
    <div className="app">
      <header className="app-header">
        <h1>TODO</h1>
      </header>

      <div className="todo-list">
        {todos.map((todo) =>
          <Todo
            key={todo.title + todo.description}
            title={todo.title}
            description={todo.description}
          />
        )}
      </div>

      <h2>Add a Todo</h2>
      <form>
        <input placeholder="Title" name="title" autoFocus={true} value={title} onChange={(val) => setTitle(val.target.value)} />
        <input placeholder="Description" name="description" value={description} onChange={(val) => setDescription(val.target.value)}/>
        <button onClick={submitToDo}>Add Todo</button>
      </form>
    </div>
  );
}

export default App;
