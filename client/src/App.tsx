import { Box, List, ThemeIcon } from "@mantine/core";
import useSWR from "swr";
import AddTodo from "./components/AddTodo";
import { CheckCircleFillIcon } from "@primer/octicons-react";

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

  async function markDone(id: number) {
    const updated = await fetch(`${API_URL}todos/${id}/done`, {
      method: "PATCH",
    }).then((res) => res.json());
    mutate(updated);
  }
  return (
    <>
      <div>
        <Box
          sx={(theme) => {
            return {
              width: "100%",
              maxWidth: "40 rem",
              margin: "0 auto",
              padding: "2rem",
            };
          }}
        >
          <List spacing="xs" size="sm" mb={12} center>
            {data?.map((todo) => {
              return (
                <List.Item
                  onClick={() => markDone(todo.id)}
                  key={`todo_list__${todo.id}`}
                  icon={
                    todo.done ? (
                      <ThemeIcon color="teal" size={24} radius="xl">
                        <CheckCircleFillIcon />
                      </ThemeIcon>
                    ) : (
                      <ThemeIcon color="grey" size={24} radius="xl">
                        <CheckCircleFillIcon />
                      </ThemeIcon>
                    )
                  }
                >
                  {todo.title}
                </List.Item>
              );
            })}
          </List>
          <AddTodo mutate={mutate} />
        </Box>
      </div>
    </>
  );
}

export default App;
