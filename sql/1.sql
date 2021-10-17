SELECT date_format(date, '%Y-%m-%d'), sum(score > 0) as num_pos_score, sum(score < 0) as num_neg_score
FROM assessments
WHERE CAST(date AS DATE) BETWEEN '2011-03-01' AND '2011-04-30'
GROUP BY date_format(date, '%Y-%m-%d');