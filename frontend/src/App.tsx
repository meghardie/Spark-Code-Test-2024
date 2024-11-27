import React, { useEffect, useState } from 'react';
import axios from 'axios';
import './App.css';
import Todo, { TodoType } from './Todo';

function App() {
  const [todos, setTodos] = useState<TodoType[]>([]);
  const [formData, setFormData] = useState<TodoType>({title: '', description: ''});

  const fetchTodos = async () => {
    try {
      const todos = await axios.get('http://localhost:8080/');
      if (todos.status !== 200) {
        console.log('Error fetching data');
        return;
      }
      else{
        setTodos(todos.data);
      }
    } catch (e) {
      console.log('Could not connect to server. Ensure it is running. ' + e);
    }
  }

  // Initially fetch todo
  useEffect(() => {
    fetchTodos()
  }, []);

  const handleSubmit = async(e: React.FormEvent) =>{
    e.preventDefault();
    try{
      let response = await axios.post("http://localhost:8080/", formData)
      console.log("Item added successfully")
      // Update list after adding new item
      fetchTodos()
    }
    catch (error){
      console.log("Error adding item")
    }
  }

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({...formData, [e.currentTarget.name]: e.currentTarget.value})
  }

  return (
    <div className="app">
      <header className="app-header">
        <h1>TODO</h1>
      </header>

      <div className="todo-list">
        {todos?.map((todo) =>
          <Todo
            key={todo.title + todo.description}
            title={todo.title}
            description={todo.description}
          />
        )}
      </div>

      <h2>Add a Todo</h2>
      <form onSubmit={handleSubmit}>
        <input placeholder="Title" name="title" autoFocus={true} onChange={handleChange}/>
        <input placeholder="Description" name="description" onChange={handleChange}/>
        <button>Add Todo</button>
      </form>
    </div>
  );
}

export default App;
