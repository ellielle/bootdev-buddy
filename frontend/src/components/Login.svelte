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
      <div>
        You aren't currently logged in! You will only have limited
        functionality.
      </div>
      <div>
        Please
        <a
          href="https://www.boot.dev/cli/login?redirect=/cli/login"
          target="_blank">click here</a
        >
        to login.
      </div>
      The login instructions can&nbsp;<a
        href="https://github.com/ellielle/bootdev-buddy#logging-in"
        target="_blank"
      >
        be found here</a
      >.
    </div>
    <div class="menu-item">
      <input
        type="text"
        placeholder="Boot.Dev CLI Code"
        bind:value={otpField}
      />
      <button on:click={loginUser}>Sign in</button>
    </div>
  </section>
</main>

<style>
</style>
