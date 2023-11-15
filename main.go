package main

import (
	"fmt"
	"math"
	"strconv"
)

// Structure pour représenter les scores des équipes
type Equipe struct {
	Nom        string
	Scores     []int
	PlusFaible int
	ScoreMoyen float64
	PlusGrand  int
}

// Fonction pour calculer le plus faible score d'une équipe
func calculerPlusFaible(equipe *Equipe) {
	plusFaible := math.MaxInt

	for _, score := range equipe.Scores {
		if score < plusFaible {
			plusFaible = score
		}
	}

	equipe.PlusFaible = plusFaible
}

// Fonction pour calculer le score moyen d'une équipe
func calculerScoreMoyen(equipe *Equipe) {
	total := 0

	for _, score := range equipe.Scores {
		total += score
	}

	equipe.ScoreMoyen = (float64(total) / float64(len(equipe.Scores))) * 2
}

// Fonction pour calculer le plus grand score d'une équipe
func calculerPlusGrand(equipe *Equipe) {
	plusGrand := math.MinInt

	for _, score := range equipe.Scores {
		if score > plusGrand {
			plusGrand = score
		}
	}

	equipe.PlusGrand = plusGrand
}

// Fonction pour saisir les résultats des matchs pour chaque équipe
func saisirResultats(equipeA, equipeB *Equipe, match int) {
	fmt.Printf("Match %d :\n", match)
	fmt.Printf("Résultats pour l'équipe %s : ", equipeA.Nom)
	equipeA.Scores[match-1] = saisirScore()
	fmt.Printf("Résultats pour l'équipe %s : ", equipeB.Nom)
	equipeB.Scores[match-1] = saisirScore()
}

// Fonction pour saisir un score en gérant les erreurs de saisie
func saisirScore() int {
	for {
		var input string
		fmt.Scan(&input)

		score, err := strconv.Atoi(input)
		if err == nil && score >= 0 {
			return score
		}

		fmt.Println("Veuillez saisir un nombre entier valide.")
	}
}

// Fonction pour afficher les résultats
func afficherResultats(equipeA, equipeB Equipe) {
	calculerPlusFaible(&equipeA)
	calculerScoreMoyen(&equipeA)
	calculerPlusGrand(&equipeA)
	calculerPlusFaible(&equipeB)
	calculerScoreMoyen(&equipeB)
	calculerPlusGrand(&equipeB)
	
	fmt.Printf("Les deux équipes\tScore match 1\tScore match 2\tScore match 3\tScore match 4\tScore match 5\tTOTAL\n")
	fmt.Printf("%s\t%d\t%d\t%d\t%d\t%d\t%d\n", equipeA.Nom, equipeA.Scores[0], equipeA.Scores[1], equipeA.Scores[2], equipeA.Scores[3], equipeA.Scores[4], equipeA.PlusFaible+equipeB.PlusFaible)
	fmt.Printf("%s\t%d\t%d\t%d\t%d\t%d\t%d\n", equipeB.Nom, equipeB.Scores[0], equipeB.Scores[1], equipeB.Scores[2], equipeB.Scores[3], equipeB.Scores[4], equipeA.PlusFaible+equipeB.PlusFaible)
	fmt.Printf("TOTAL\t\t%d\t%d\t%d\t%d\t%d\t%d\n", equipeA.Scores[0], equipeA.Scores[1], equipeA.Scores[2], equipeA.Scores[3], equipeA.Scores[4], equipeA.PlusFaible+equipeB.PlusFaible)

	fmt.Printf("\n\t\t\t\tSYNTHESE\n")
	fmt.Printf("Equipes\tPlus faible score\tScore moyen\tPlus grand score\n")
	fmt.Printf("%s\t%d\t\t    %.1f\t\t\t%d\n", equipeA.Nom, equipeA.PlusFaible, equipeA.ScoreMoyen, equipeA.PlusGrand)
	fmt.Printf("%s\t%d\t\t    %.1f\t\t\t%d\n", equipeB.Nom, equipeB.PlusFaible, equipeB.ScoreMoyen, equipeB.PlusGrand)
	fmt.Printf("TOTAL\t\t%d\t\t    %.1f\t\t\t%d\n", equipeA.PlusFaible+equipeB.PlusFaible, (equipeA.ScoreMoyen+equipeB.ScoreMoyen)/2, equipeA.PlusGrand+equipeB.PlusGrand)
}

func main() {
	// Définir les équipes
	equipeA := Equipe{Nom: "Equipe A", Scores: make([]int, 5)}
	equipeB := Equipe{Nom: "Equipe B", Scores: make([]int, 5)}

	for {
		// Menu interactif
		fmt.Println("\nMenu :")
		fmt.Println("1. Saisir les résultats des matchs")
		fmt.Println("2. Afficher les résultats")
		fmt.Println("3. Quitter")

		var choix int
		fmt.Print("Votre choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			// Saisir les résultats des matchs pour chaque équipe
			for i := 1; i <= 5; i++ {
				saisirResultats(&equipeA, &equipeB, i)
			}
		case 2:
			// Afficher les résultats
			afficherResultats(equipeA, equipeB)
		case 3:
			// Quitter le programme
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide. Veuillez choisir une option valide.")
		}
	}
}
