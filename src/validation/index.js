  import Schema from "async-validator";
  import login from './constraint/login/index.js';

  let allValidation = {};

  export function newValidator(descriptor = {}, translate = () => {}) {
    let validator = new Schema(descriptor, {
      messages: translate("asyncValidator"),
    });
    validator._messages = translate("asyncValidator");
    return validator
  }

  // load all constraints with translate
  export function loadValidation(translate) {
    console.log(login(translate))
    allValidation = {
      constraints: {
        login: login(translate),
      }
    }
  }

  export function getValidation() {
    return allValidation
  }