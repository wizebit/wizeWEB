const checkValidity = (value, rules, fieldName = null) => {
    let validationObj = {
        isValid: true,
        errorMessage: null
    };

    const name = fieldName ? fieldName.replace(/([A-Z])/g, ' $1').replace(/^./, str => str.toUpperCase()) : null;
    if (rules) {
        if (rules.required) {
            validationObj.isValid = value.trim() !== '' && validationObj.isValid;
            validationObj.errorMessage = fieldName ? `${name} is required` : 'This field is required';
        }

        if (rules.minLength) {
            validationObj.isValid = value.length >= rules.minLength && validationObj.isValid;
            validationObj.errorMessage = fieldName ? `${name} value is shorter than required` : 'This field value is shorter than required';
        }

        if (rules.maxLength) {
            validationObj.isValid = value.length <= rules.maxLength && validationObj.isValid;
            validationObj.errorMessage = fieldName ? `${name} value is longer than required` : 'This field value is longer than required';
        }

        if (rules.isEmail) {
            const pattern = /[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?/;
            validationObj.isValid = pattern.test(value) && validationObj.isValid;
            validationObj.errorMessage = fieldName ? `${name} is incorrect` : 'Incorrect email';
        }

        if (rules.isPhone) {
            const pattern = /^[+]?[(]?[0-9]{3}[)]?[-\s.]?[0-9]{3}[-\s.]?[0-9]{4,6}$/;
            validationObj.isValid = pattern.test(value) && validationObj.isValid;
            validationObj.errorMessage = fieldName ? `${name} is incorrect` : 'Incorrect phone';
        }

        if (rules.isNumeric) {
            const pattern = /^\d+$/;
            validationObj.isValid = pattern.test(value) && validationObj.isValid;
            validationObj.errorMessage = fieldName ? `${name} is incorrect` : 'This field is incorrect';
        }
    }

    return validationObj;
};

export default checkValidity;