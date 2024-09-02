<svelte:head>
	<title>TODO List</title>
	<meta name="description" content="Main page" />
</svelte:head>

<script lang="ts">
    import { createForm } from 'felte';
    import { onMount } from 'svelte';
    import { goto } from "$app/navigation";
    import ky from 'ky';
    import { Circle} from 'svelte-loading-spinners';
	import { navigating } from '$app/stores'
    
    interface Task {
        ID: string;
        UserId: string;
        title: string;
        description: string;
        Due_date: string;
    }

    let para = 'Todo app';
    let tasks: any[] = [];

    let token: string | null = null;
    $: errorString = "";

    onMount(() => {
    token = localStorage.getItem('token');
    console.log('Retrieved Token:', token);
    if (!token) {
        goto('/');
        return;
    }
    getTasks();
});
    
//     if (browser) {
//     token = localStorage.getItem('token');
//     if (!token) {
//       goto('/');
//     }
//   }
    
//   if (!token && browser) {
//   goto('/');
// }

    const { form, data, errors, touched } = createForm({
        initialValues:{
            title: '',
            description: '',
            due_date:'',
            // status: '',
        },
        onSubmit: async (values) => {

            if (values.due_date) {
            const date = new Date(values.due_date);
            values.due_date = date.toISOString(); // Convert to ISO string
            }
            const newTask = await ky.post("http://192.168.18.120:8000/todo", {json: values, headers: {'Authorization': `Bearer ${token}`}}).json();
            tasks =[...tasks,newTask];
            console.log("Task created: ", newTask);
            // console.log('Submitting task:', values); // Log the task values being submitted
            // try {
            // const response = await fetch('http://192.168.18.120:8000/todo', {
            // method: 'POST',
            // headers: {
            //     'Content-Type': 'application/json',
            //     'Authorization': `Bearer ${token}`
            // },
            // body: JSON.stringify(values), // Ensure this matches the expected format
            // });
            //     console.log('Task creation response status:', response.status);
            //     if (response.ok) {
            //         console.log('Task creation in progress');
            //         const newTask = await response.json();
            //         tasks = [...tasks, newTask];
            //     } else {
            //         console.error('Failed to add task');
            //     }
            // } catch (error) {
            //     console.error('Error adding task:', error);
            // }
        },
        onError : (errors) => {
            console.log(errors);
            errorString = errors as string;
        }
    });

    async function getTasks() {
        const fetchedTasks = await ky.get("http://192.168.18.120:8000/todo", {headers: {'Authorization': `Bearer ${token}`}}).json<Task[]>();
        console.log("Tasks Fetched: ", fetchedTasks);
        tasks = [...fetchedTasks];
        // try {
        //     const response = await fetch('http://192.168.18.120:8000/todo', {
        //         //method: 'GET',
        //         headers: {
        //             'Authorization': `Bearer ${token}`
        //         }
        //     });
        //     if (response.ok) {
        //         const fetchedTasks = await response.json();
        //         console.log('Tasks fetched:', fetchedTasks); // Log fetched tasks
        //         tasks = [...fetchedTasks]; // Assign fetched tasks to the tasks 
        //         console.log('Updated tasks array:', tasks); 
        //     } else {
        //         console.error('Failed to fetch tasks');
        //     }
        // } catch (error) {
        //     console.error('Error fetching tasks', error);
        // }
        // console.log('Rendering tasks:', tasks);
    }

    async function deleteTask(id: any) {
        await ky.delete(`http://192.168.18.120:8000/todo/${id}`, {headers:{'Authorization': `Bearer ${token}`}});
        tasks = tasks.filter(task => task.id !== id);
        console.log("Task deleted: ", tasks);
        // try {
        //     const response = await fetch(`http://192.168.18.120:8000/todo/${id}`, {
        //         method: 'DELETE',
        //         headers: {
        //             'Authorization': `Bearer ${token}`
        //         }
        //     });
        //     if (response.ok) {
        //         tasks = tasks.filter(task => task.id !== id);
        //     } else {
        //         console.error('Failed to delete task');
        //     }
        // } catch (error) {
        //     console.error('Error deleting task', error);
        // }
    }

    function toggleComplete(id: any) {
    tasks = tasks.map(task =>
        task.ID === id ? { ...task, completed: !task.completed } : task
    );
    }
//     onMount(() => {
//     if (!token) {
//       goto('/');
//       return;
//     }
//       getTasks();
// });
</script>

{#if $navigating}
	<Circle size="60" color="#FF3E00" unit="px" duration="1s" />
{/if}

<section>
	<header>
		<h1>{para.toUpperCase()}</h1>
	</header>
	<h2>
		<strong>Add tasks todo</strong>
	</h2>
	<form use:form>
	<section>
		<div class="input">
		<label for="title">Task Title</label>
		<input type="text" placeholder="Task title" name="title" bind:value={data.title}>
		<label for="description">Task Description</label>
		<input type="text" placeholder="Task description" name="description" bind:value={data.description}>
		<label for="due_date">Due Date/time</label>
		<input type="datetime-local" id="duedate" name="due_date" bind:value={data.due_date}> 
        </section>
		<section class="button">
			<button type="submit">
				Add Task
			</button>
			<button type="button" on:click={getTasks}>
				Refresh tasks
			</button>
		</section>
	</form>
	<section class="task-container">
        <h2>Tasks</h2>
        <div class="task-list">
            {#each tasks as task (task.ID)}
                <div class="task">
                    <div class="task-header">
                        <input type="checkbox" checked={task.completed} on:change={() => toggleComplete(task.ID)}>
                        <span class:completed={task.completed} class="task-title">{task.title}</span>
                    </div>
                    <p class:completed={task.completed} class="task-description">{task.description}</p>
                    <p class:completed={task.completed} class="task-due-date">Due: {new Date(task.Due_date).toLocaleString()}</p>
                    <button on:click={() => deleteTask(task.ID)}>Delete</button>
                </div>
            {/each}
        </div>
    </section>
	<footer>
		<p>This is a shit design but bear with me</p>
	</footer>
</section>

<style>
    section {
        max-width: 600px;
        margin: 0 auto;
        padding: 20px;
        font-family: Arial, sans-serif;
    }
    
    h1, h2 {
        text-align: center;
        font-size: xx-large;
    }
    
    .input {
        display: flex;
        flex-direction: column;
        align-items: center;
        margin-bottom: 20px;
        font-size: xx-large;
        width: 100%;
    }
    
    .input label {
        margin-bottom: 5px;
        text-align: center; /* Center the label text */
        width: 100%;
    }
    
    .input input {
        padding: 8px;
        margin-bottom: 10px;
        width: 350px;
    }
    
    .button {
        display: flex;
        justify-content: center;
        gap: 10px;
    }
    
    button {
        padding: 10px 20px;
        background-color: #4CAF50;
        color:antiquewhite;
        border: none;
    }
    
    .task-container {
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    .task-list {
        display: flex;
        flex-wrap: wrap;
        justify-content: center;
        gap: 1rem;
    }
    
    .task {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    background-color:azure;
    padding: 1rem;
    border-radius: 0.5rem;
    box-shadow: black;
    width: 200px;
    }
    .task-header {
    display: flex;
    align-items: center;
    }

    .task-title {
    font-weight: bold;
    margin-bottom: 0.5rem;
    }

    .task-description{
    margin-bottom: 0.5rem;
    }

   .task-due-date {
    font-size: 0.9rem;
    color: #666;
    }
    
    .task span {
        flex-grow: 1;
    }
    
    .task .completed {
        text-decoration: line-through;
        color: gray;
    }
    
    footer {
        margin-top: 30px;
        text-align: center;
        color: gray;
    }
    .completed {
        text-decoration: line-through;
        color: gray;
    }
</style>