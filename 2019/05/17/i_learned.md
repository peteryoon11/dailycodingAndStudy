http://webframeworks.kr/tutorials/expressjs/expressjs_orm_one/

검색어 sequelize 사용법




https://stackoverflow.com/questions/14169655/sequelize-js-foreign-key


var Person = sequelize.define('Person', {

        name: Sequelize.STRING
});

var Father = sequelize.define('Father', {

        age: Sequelize.STRING,
        //The magic start here
        personId: {
              type: Sequelize.INTEGER,
              references: 'persons', // <<< Note, its table's name, not object name
              referencesKey: 'id' // <<< Note, its a column name
        }
});

Person.hasMany(Father); // Set one to many relationship


sequelize.define
