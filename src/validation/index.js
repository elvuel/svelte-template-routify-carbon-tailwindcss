  import Schema from "async-validator";

  export function newValidator(descriptor = {}, translate = () => {}) {
    let validator = new Schema(descriptor, {
      messages: translate("asyncValidator"),
    });
    validator._messages = translate("asyncValidator");
    return validator
  }