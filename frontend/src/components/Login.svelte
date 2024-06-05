<script>
  import { LoginUserWithOTP } from "../../wailsjs/go/main/App.js";
  export let loggedIn;
  // Empty field initially for the one-time password
  let otpField = "";

  // loginUser takes a user's OTP, trades it for an access token which
  // is saved for futher use, and marks the user as logged in
  function loginUser() {
    LoginUserWithOTP(otpField).then((result) => (loggedIn = result));
  }
</script>

<main>
  <section>
    <div class="menu-item btn-login">
      <p>
        You aren't currently logged in! You will only have limited
        functionality.
      </p>
      <p>
        Please
        <a
          href="https://www.boot.dev/cli/login?redirect=/cli/login"
          target="_blank"
          class="text-primary-500">click here</a
        >
        to get your one-time password.
      </p>

      <div>
        The login instructions can&nbsp;<a
          href="https://github.com/ellielle/bootdev-buddy#logging-in"
          target="_blank"
          class="text-primary-500"
        >
          be found here</a
        >.
      </div>
    </div>
    <div class="menu-item mt-8">
      <input
        type="text"
        placeholder="Boot.Dev CLI Code"
        bind:value={otpField}
      />
      <button
        class="text-primary-500 border rounded px-2 py-1.5"
        on:click={loginUser}>Sign in</button
      >
    </div>
  </section>
</main>

<style>
  .menu-item > p {
    margin-bottom: 0.5rem;
  }
</style>
