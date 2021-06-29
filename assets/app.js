// Sign in function
async function do_auth(event) {
  event.preventDefault();
  hideAuthError();
  const authInitResponse = await fetch('/authentication_initialize', {
    method: 'GET',
  });

  if (!authInitResponse.ok) {
    // TODO: showError
    console.log((await authInitResponse.json()).error);
    showAuthError(await authInitResponse.json().error);
  }

  const authOptions = await authInitResponse.json();

  const authenticatorResponse = await hankoWebAuthn.get(authOptions);

  const authenticationResponse = await fetch('/authentication_finalize', {
    method: 'POST',
    body: JSON.stringify(authenticatorResponse),
  });

  if (!authenticationResponse.ok) {
    // TODO: show Error
    console.log((await authenticationResponse.json()).error);
    showAuthError(await authInitResponse.json().error);
  } else {
    location.assign('/content');
  }
}

async function do_reg(event) {
  event.preventDefault();
  const username = document.getElementById('username').value;
  let query = '?user_name=' + username;
  const regInitResponse = await fetch('/registration_initialize' + query, {
    method: 'GET',
  });

  if (!regInitResponse.ok) {
    const error = (await regInitResponse.json()).error;
    showRegError(error);
  }

  const creationOptions = await regInitResponse.json();

  const authenticatorResponse = await hankoWebAuthn.create(creationOptions);

  const registrationResponse = await fetch('/registration_finalize', {
    method: 'POST',
    body: JSON.stringify(authenticatorResponse),
  });

  if (!registrationResponse.ok) {
    const error = (await registrationResponse.json()).error;
    showError(error);
  } else {
    location.assign('/content');
  }
}

function showRegError(error) {
  console.log(error)
  document.getElementById("username").classList.add("is-invalid")
  document.getElementById("invalidRegistrationFeedback").innerHTML = error
}

const showAuthError = (error) => {
  console.error('Error: ', error);
  document.querySelector('.alertMessage').textContent = error;
  document.querySelector('.alertBox').classList.remove('alertBox--hidden');
};

const hideAuthError = () => {
  document.querySelector('.alertMessage').textContent = '';
  document.querySelector('.alertBox').classList.add('alertBox--hidden');
};