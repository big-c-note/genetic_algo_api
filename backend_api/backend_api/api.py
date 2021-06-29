from functools import partial
from typing import List

from flask import Flask, jsonify
import requests

from backend_api.gen_algo import (
    generate_population,
    run_evolution,
    Thing,
    Genome,
    fitness,
)


app = Flask(__name__)
app.config["DEBUG"] = True


@app.route("/", methods=["GET"])
def home():
    return """<h1>Genetic Algo API</h1>"""


@app.route("/run_algo", methods=["GET"])
def run_algo():
    host = "http://localhost:8080"
    r = requests.get(host + "/things")
    things = r.json()

    for i, thing in enumerate(things):
        things[i] = Thing(**thing)

    populations, generations = run_evolution(
        population_func=partial(
            generate_population, size=10, genome_length=len(things)
        ),
        fitness_func=partial(fitness, things=things, weight_limit=500),
        fitness_limit=2000,
        generation_limit=100,
    )

    def genome_to_things(genome: Genome, things: List[Thing]) -> List[Thing]:
        result = []
        for i, thing in enumerate(things):
            if genome[i] == 1:
                result += [thing.name]
        return result

    result = {}
    result["number of generations"] = str(generations)
    result["best solution"] = genome_to_things(populations[0], things)

    return jsonify(result), 201


if __name__ == "__main__":
    app.run(port=8083)
