import { Box } from "@mantine/core";
import useSWR from "swr";
import AddTodo from "./components/AddTodo";

export const API_URL = "http://localhost:5000/api/";
export interface Todo {
  id: number;
  title: string;
  body: string;
  done: boolean;
}

const fetcher = (url: string) =>
  fetch(`${API_URL}${url}`).then((res) => res.json());

function App() {
  const { data, mutate } = useSWR<Todo[]>("todos", fetcher);
  return (
    <>
      <div>
        <Box>
          {JSON.stringify(data)}
          <AddTodo mutate={mutate} />
        </Box>
      </div>
    </>
  );
}

export default App;
