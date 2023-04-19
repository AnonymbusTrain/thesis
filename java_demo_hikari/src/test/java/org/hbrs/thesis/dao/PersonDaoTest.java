package org.hbrs.thesis.dao;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;

import java.sql.SQLException;
import java.util.List;

import org.hbrs.thesis.model.Person;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.Test;

class PersonDaoTest {
    private PersonDao personDao = new PersonDao();

    @AfterEach
    void teardown() {
        try {
            personDao.dropDBTable();
        } catch (SQLException ex) {
            ex.printStackTrace();
        }
    }

    @Test
    void assureThatThousandPersonsAreGenerated() throws SQLException {
        long numberOfPersonsToGenerate = 1000;
        personDao.generateNumberOfRandomPersonsToDB(1000);
        assertEquals(numberOfPersonsToGenerate, personDao.getAllPersons().size());
    }

    @Test
    void assureThatPersonsGeneratedCorrectly() throws SQLException {
        long numberOfPersonsToGenerate = 100;
        personDao.generateNumberOfRandomPersonsToDB(numberOfPersonsToGenerate);
        List<Person> persons = personDao.getAllPersons();
        for (Person person : persons) {
            assertFalse(person.getFirstName().isBlank());
            assertFalse(person.getLastName().isBlank());
            assertFalse(person.getBirthDate().toString().isBlank());
            assertFalse(person.getTimestamp().toString().isBlank());
            long id = person.getId();
            assertFalse(id < 0 && id > numberOfPersonsToGenerate);
        }
    }
}