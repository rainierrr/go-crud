import * as React from "react";
import AddIcon from '@material-ui/icons/Add'
import Card from '@material-ui/core/Card'
import CardContent from '@material-ui/core/CardContent'
import Fab from '@material-ui/core/Fab'
import axios from 'axios';

const  test = async() => {
  const url = "http://localhost:8080/tasks"
  try {
    const res = await axios.get(url);
    console.log(res);
  } catch (error) {
    console.log(error);
  }
}

const TaskCard: React.FC = () => {
  test()
  return (
    <Card>
    <CardContent>
      <h1>課題一覧</h1>
      <h2>今日やること</h2>
      <Fab size="small" color="secondary">
        <AddIcon />
      </Fab>
    </CardContent>
  </Card>
  )
}

export default TaskCard
