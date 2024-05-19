import { createLazyFileRoute } from '@tanstack/react-router';
import DisplayTasks from '../components/DisplayTasks';

export const Route = createLazyFileRoute('/')({
  component: Index,
});

function Index() {
  return (
    <div>
      <h3>Welcome Home!</h3>
      <DisplayTasks />
    </div>
  );
}
