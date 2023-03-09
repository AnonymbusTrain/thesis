package org.hbrs.thesis;

import static spark.Spark.port;
import static spark.Spark.exception;

import org.hbrs.thesis.config.ApplicationConfig;
import org.hbrs.thesis.controller.MetricsController;
import org.hbrs.thesis.controller.PersonController;
import org.hbrs.thesis.dto.MessageDto;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import com.google.gson.Gson;
public final class App {
    private static final Logger logger = LoggerFactory.getLogger(App.class);

    private App() {
    }

    /**
     * Generates 3 Person endpoints
     * @param args The arguments of the program.
     */
    public static void main(String[] args) {
        ApplicationConfig applicationConfig = new ApplicationConfig();
        port(applicationConfig.getApplicationPort());

        PersonController personController = new PersonController();
        personController.generatePersons();
        personController.getAllPersons();
        personController.removeTable();
        personController.randomCaclulation();
        MetricsController metricsController = new MetricsController();
        metricsController.exposePrometheusMetrics();
        Gson gson = new Gson();
        exception(Exception.class, (exception, req, res) -> {
            res.type("application/json");
            res.status(400);
            MessageDto messageDto = new MessageDto("Something went wrong");
            res.body(gson.toJson(messageDto));
            logger.warn("There was an exception: {}, that was thrown with the message: {}", exception.getClass().getSimpleName(), exception.getLocalizedMessage());
        });

    }
}
