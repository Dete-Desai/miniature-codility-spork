//1.IMPORT AND EXPORT
// import {apiKey} from "./util.js";

//import apiKey from "./util.js";
import { apiKey, abc as content } from "./util.js";
// import * as util from "./util.js";

//console.log(util.apiKey)
console.log(content)

//2. VARIABLES AND VALUES

let userMessage = "Hello World";
userMessage = "New Value";
console.log(userMessage)

const userMessage1 = "Hello World";
console.log(userMessage1);

//3. OPERATORS
console.log(10 + 10)
console.log(10 - 10)
console.log(10 * 10)
console.log(10 / 10)
console.log(10 === 10)
console.log(10 > 10)
console.log(10 < 10)
console.log(10 >= 10)
console.log(10 <= 10)

//4. FUNCTIONS AND PARAMETERS
function createGreeting(userName, message = "Hello"){
    return "Hi, I am" + userName + ". " + message
}

const greeting1 = createGreeting("Max");
console.log(greeting1);

const greeting2 = createGreeting("Manuel", "Hello, what's up?");
console.log(greeting12);

//5. ARROW FUNCTIONS
export default (userName, message)=> {
    console.log("Hello"); 
    return userName + message;
}

number => ({ age: number });

//6. OBJECTS AND CLASSES
//Step1
const user = {
    name: "Max",
    age: 34,
    greet(){
        console.log("Hello")
        console.log(this.age);
    }
};
console.log(user.name);
user.greet();

//Step2
class User {
    constructor(name, age){
        this.name = name;
        this.age = age;
    }
    greet(){
        console.log('Hi');
    }
}

const user1 = new User("Manuel", 35);
console.log(user1)
user1.greet;

//ARRAY AND ARRAYS METHODS
//Arrays $ Array Methods like map()

const hobbies = ["Sports","Cooking","Reading"];
console.log(hobbies[0])

hobbies.push("Working")
console.log(hobbies)

//Step 1
const index = hobbies.findIndex((item)=>{
    return item === "Sports";
});

//step 2
const index1 = hobbies.findIndex((item)=> item === "Cooking"); 

console.log(index)

const editedHobbies = hobbies.map((item) => item + "!");

console.log(editedHobbies)

const editedHobbies1 = hobbies.map((item)=>({text: item}));

console.log(editedHobbies1)

//DESTRUCTERING
const [firstName, lastName]=["Faith","Suyianka"];
console.log(firstName)
console.log(lastName)

const {name: userName, age} = {
    name: "Lamech",
    age: 34
}

console.log(userName);
console.log(age);

//The destructuring syntax explained in the previous 
//lecture can also be used in function parameter lists.
// Here's an example:

// function storeOrder(order) {
//   localStorage.setItem('id', order.id);
//   localStorage.setItem('currency', order.currency);
// }

function storeOrder({id, currency}) { // destructuring
    localStorage.setItem('id', id);
    localStorage.setItem('currency', currency);
  }

storeOrder({id: 5, currency: 'USD', amount: 15.99});

console.log(storeOrder.id);

//THE SPREAD OPERATOR
const skills = ["Drawing, Scoring"];
const user2 = {
    name: "Halaand",
    age: 22
}; 

const newSkills = ["Reading"];

const mergedSkills = [...skills, ...newSkills];
console.log(mergedSkills);

const extendedUser = {
    isAdmin: true, 
    ...user
};

console.log(extendedUser);

//CONTROL STRUCTURES
//If Statement
const password = prompt('Your password');

if (password === "Hello") {
    console.log("Hello works");
} else if (password === "hello") {
    console.log("hello works");
}  else {
    console.log("Access not granted");
}

//For Loop Statement
for (const hobby of hobbies){
    console.log(hobby);
}

const list = document.querySelector("ul");
list.remove();

//USING FUNCTIONS AS VALUES

function handleTimeOut(){
    console.log("Timed out");
}

const handleTimeOut2 = () =>{
    console.log("Timed out ... again!");
};

setTimeout(handleTimeOut, 2000);
setTimeout(handleTimeOut2, 3000);

setTimeout(()=>{
    console.log("More timing out...");
}, 4000);

function greeter(greetFn){
    greetFn();
}

greeter(()=> console.log("Hi"));

//DEFINING FUNCTIONS INSIDE A FUNCTION
function init(){
    function greet(){
        console.log('Hi');
    }

    greet();
}

init();

//REFERENCE VS PRIMITIVE VALUES
let userMessage2 = 'Hello';
userMessage2 = userMessage2.concat('!!!');
hobbies.push("Working");
console.log(hobbies);

