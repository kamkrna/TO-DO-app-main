<script lang="ts">
    import { goto } from "$app/navigation";
	import { createForm } from 'felte'
    import { error } from "@sveltejs/kit";
    import ky from "ky";
    $: errorString = ""
	let para = 'Todo app';	
	const { form, errors, isSubmitting } = createForm({
    initialValues: {
      username: '',
      email: '',
      password: '',
    },
    onSubmit: async (values) => {
      const json = await ky.post("http://192.168.18.120:8000/registration", {json: values}).json();
      goto("/");
      }, 
      onError: (errors) => {
      console.log(errors)
      errorString = errors as string
    }
    },
  );

    // function setErrors(arg0: { form: string; }) {
    //     throw new Error("Function not implemented.");
    // }
</script>

<svelte:head>
	<title>Login</title>
	<meta name="description" content="User Registration" />
</svelte:head>

<section>
	<header>
		<h1>{para.toUpperCase()}</h1>
	</header>
	
	<h2>
		User Registration
	</h2>

	<h2>
		Enter your username and password and email.
	</h2>

	<form use:form>
		<div class="input">
		<input type="text" name="username" placeholder="Username">
		<input type="password" name="password" placeholder="Password">
        <input type="email" name="email" placeholder="Email">
	</div>
		<div class="button">
			<button style="margin-left: 20px;" type="submit">
				Register User
			</button>
		</div>
	</form>
    {#if error}
    <p class="err">{error}</p>
    {/if}
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
	
	h2 {
        font-size: xx-large;
        text-align: center;
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
