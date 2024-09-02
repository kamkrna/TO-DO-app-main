<svelte:head>
	<title>Login</title>
	<meta name="description" content="Login page" />
</svelte:head>

<script lang="ts">
	import { goto } from "$app/navigation";
	import { createForm } from 'felte'
    import ky from "ky";
    import { Jellyfish } from "svelte-loading-spinners";

    $: errorString = "";
	let loading = false;

	const { form } = createForm({
    initialValues: {
      username: '',
      password: '',
    },
    onSubmit: async (values) => {
		loading = true;
      const res: {token: string} = await ky.post("http://192.168.18.120:8000", { json: values}).json()

      localStorage.setItem('token', res.token)

      goto("/todo")
    },
    onError: (errors) => {
      console.log(errors)
      errorString = errors as string
    }
  });
	let para = 'Todo app';
    

    const redirectRegister = () => goto('/registration');
</script>

{#if loading}
	<div class="container">
		Loading please wait...
		<div class="loader"><Jellyfish size="60" color="#FF3E00" unit="px" duration="1s" /></div>
	</div>
{:else}
<section>
	<header>
		<h1>{para.toUpperCase()}</h1>
	</header>
	
	<h2>
		<strong>Log In</strong>
	</h2>

	<h2>
		Enter your username and password.
	</h2>

	<form use:form>
		<div class="input">
		<input type="text" placeholder="Username" name="username">
		<input type="password" placeholder="Password" name="password">
	</div>
		<div class="button">
			<button type="submit">
				Login
			</button>
			<button style="margin-left: 20px;" type="button" on:click={redirectRegister}>
				Register
			</button>
		</div>
		{#if errorString}
  		<p class="err" style="font-size :larger;">{errorString}</p>
		{/if}
	</form>
	<footer>
		<p>This is a shit design but bear with me</p>
	</footer>
</section>
{/if}

<style>
	.container {
        margin-top: 50%;
	}

	.loader{
		margin-left: 40%;
	}
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

	footer{
		margin-top: 30px;
        text-align: center;
        color: gray;
	}
	
	div{
		font-size: xx-large;
		font-family: 'Times New Roman', Times, serif;
		font-style: italic;
		font-weight: bolder;
	}
	.input{
        display: flex;
        flex-direction: column;
        align-items: center;
        margin-bottom: 20px;
        width: 100%;
    }
    .input input {
        padding: 8px;
        margin-bottom: 10px;
        width: 350px;
    }
    .button{
        display: flex;
        justify-content: center;
        gap: 10px;   
    }
    button{
        padding: 10px 20px;
        background-color: #4CAF50;
        color: white;
        border: none;
    }
    .err{
      text-align: center;
    }
</style>
