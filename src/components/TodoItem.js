const TodoItem = ({ todo }) => {
    return (
      <li className={`${todo.done ? "checked" : ""}`}> 
      {todo.id}. Nama Depan : {todo.firstname}, Nama Belakang : {todo.lastname} <span className="close">x</span>
      </li>
    );
  };
  
  export default TodoItem;
  