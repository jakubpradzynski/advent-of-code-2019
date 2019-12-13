(ns pl.jakubpradzynski.advent-of-code-2019.day-8 
    (:require [clojure.string :as string]))

(def WIDE 25)
(def TALL 6)

(defn ReadFileToString [filepath]
    (slurp filepath)
)

(defn GetLayersCount [image wide tall]
    (/ (count image) (* wide tall))
)

(defn GetLayer [image layerNumber wide tall]
    (def firstDigitIndex (* layerNumber wide tall))
    (def lastDigitIndex (* (+ layerNumber 1) wide tall))
    (subvec image firstDigitIndex lastDigitIndex)
)

(defn GetParticularDigitsCount [layer, digit]
    (def filtered (filter #(= % digit) layer))
    (count filtered)
)

; Check if command line args contains only path to file with encoded image
(def args *command-line-args*)
(if (not= 1 (count args))
    (
        (println "You should pass filepath to encoded image!") 
        (System/exit 1)
    )
)

; Load image to vector of numbers
(def image (string/split (string/trim (ReadFileToString (first args))) #""))

(def layersCount (GetLayersCount image WIDE TALL))

(def layerWithLawestZeroDigitCount (first 
    (sort #(compare (GetParticularDigitsCount %1 "0") (GetParticularDigitsCount %2 "0")) (
        map #(GetLayer image % WIDE TALL) (range 0 layersCount)
    ))
))
(def zeroAndOneProduct (* (GetParticularDigitsCount layerWithLawestZeroDigitCount "1") (GetParticularDigitsCount layerWithLawestZeroDigitCount "2")))
(print "Part One - ")
(println zeroAndOneProduct)

(def resultLayer (GetLayer image (- layersCount 1) WIDE TALL))
(def newLayersCount (- layersCount 1))
(dotimes [i newLayersCount] 
    (def index (- (- layersCount 1) (+ i 1)))
    (def layer (GetLayer image index WIDE TALL))
    (doseq [[j digit] (map-indexed vector layer)]
        (if (not= digit "2") 
            (def resultLayer (assoc resultLayer j digit))
        )
    )
)
(println "Part Two")
(loop [x 0]
    (when (< x (* WIDE TALL))
       (println (subvec resultLayer x (+ x WIDE)))
       (recur (+ x WIDE))))